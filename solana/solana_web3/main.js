const prompt = require('prompt-async');
const yargs = require("yargs");
const web3 = require("@solana/web3.js");
const fs = require("fs");

var wallet;
var connection;

function getWallet(wallet_str) {
    try {
        data = fs.readFileSync(`./${wallet_str}.key`)
        wallet = web3.Keypair.fromSecretKey(data)
        console.log(`Loaded wallet ${wallet_str} [ ${wallet.publicKey.toString()} ]`)
        return wallet

    } catch(e) {
        console.log(`Wallet ${wallet_str} not found. Generating...`)
        wallet = web3.Keypair.generate()
        fs.writeFileSync(`./${wallet_str}.key`, wallet.secretKey)
        console.log(`Wallet ${wallet_str} [ ${wallet.publicKey.toString()} ] generated!`)
        return wallet
    }
}

function getBalance(wallet) {
    return connection.getBalance(wallet.publicKey)
}

async function createStakeAccount(main, authorised, staking, amount) {
    let createAccountTransaction = web3.StakeProgram.createAccount({
        fromPubkey: main.publicKey,
        authorized: new web3.Authorized(authorised.publicKey, authorised.publicKey),
        lamports: amount,
        lockup: new web3.Lockup(0, 0, main.publicKey),
        stakePubkey: staking.publicKey
    });
    let tx = await web3.sendAndConfirmTransaction(connection, createAccountTransaction, [main, staking]);
    console.log(tx)

}

async function requestAirdrop(wallet) {
    let airdropSignature = await connection.requestAirdrop(wallet.publicKey, web3.LAMPORTS_PER_SOL);
    console.log(await connection.confirmTransaction(airdropSignature))
    let newBalance = await getBalance(mainWallet)
    console.log(`Main ballance ${newBalance}`)
}


mainWallet = getWallet("main")
authorisedWallet = getWallet("authorised")
stakingWallet = getWallet("staking")

// establish connection to devnet
connection = new web3.Connection(web3.clusterApiUrl('devnet'), 'confirmed');

async function delegateStaking(mainWallet, stakingWallet, authorisedWallet) {

    // To delegate our stake, we get the current vote accounts and choose the first
    let voteAccounts = await connection.getVoteAccounts()
    let voteAccount = voteAccounts.current.concat(
        voteAccounts.delinquent,
    )[0];
    let votePubkey = new web3.PublicKey(voteAccount.votePubkey);

    // We can then delegate our stake to the voteAccount
    let delegateTransaction = web3.StakeProgram.delegate({
        stakePubkey: stakingWallet.publicKey,
        authorizedPubkey: authorisedWallet.publicKey,
        votePubkey: votePubkey,
    });
    await console.log(web3.sendAndConfirmTransaction(connection, delegateTransaction, [mainWallet, authorisedWallet]));

}


// Call start
(async() => {

    let fromBalance = await getBalance(mainWallet);
    console.log(`Main wallet balance: ${fromBalance} lamports`)
    if (fromBalance == 0) {
        prompt.start();
        console.log('Main wallet balance is zero. Want airdrop yes/no?');
        const result = await prompt.get(['answer'])
        if (result.answer.toLowerCase() == 'yes') {
                requestAirdrop(mainWallet)
        }
    }

    stakeAmount = fromBalance - 10000 // fee!

    let minForStakeAccount = (await connection.getMinimumBalanceForRentExemption(web3.StakeProgram.space))

    if (stakeAmount < minForStakeAccount) {
        console.log(`Not enough lamports for staking, need at least ${minForStakeAccount}`)
    }

    var stakeState
    try {
        // check if staking already active
        stakeState = await connection.getStakeActivation(stakingWallet.publicKey);
    } catch (e) {
        console.log("Staking account does not exist, creating...")
        createStakeAccount(mainWallet, authorisedWallet, stakingWallet, stakeAmount)
        stakeState = await connection.getStakeActivation(stakingWallet.publicKey);
    }
    if (stakeState) {
        console.log('Staking account state:', stakeState)
    } 
    
    console.log(`Staking account balance: ${await getBalance(stakingWallet)} lamports`)
    if (stakeState.state == 'inactive') {
        console.log('Delegating staking to first voter in the list')
        delegateStaking(mainWallet, stakingWallet, authorisedWallet)
    }

})();
