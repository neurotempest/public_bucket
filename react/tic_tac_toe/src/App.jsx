import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

export default function Game() {

  const [xIsNext, setXIsNext] = useState(true);
  const [history, setHistory] = useState([Array(9).fill(null)]);

  const currentBoardState = history[history.length-1];

  function handlePlay(newBoardState) {

    setHistory([...history, newBoardState]);
    setXIsNext(!xIsNext);
  }

  return (
    <div className="game">
      <div className="game-board">
        <Board xIsNext={xIsNext} boardState={currentBoardState} handlePlay={handlePlay}/>
      </div>
      <div className="game-history">
        <ol>{/*tofo*/}</ol>
      </div>
    </div>
  )
}

function Board({xIsNext, boardState, handlePlay}) {

  function handleClick(i) {
    return () => {

      if (boardState[i] || calculateWinner(boardState)) {
        return;
      }

      const updatedBoard = boardState.slice();

      let val = "";
      if (xIsNext) {
        val = "X";
      } else {
        val = "O";
      }

      updatedBoard[i] = val;

      handlePlay(updatedBoard);
    }
  }

  console.log("calcing winner");

  const winner = calculateWinner(boardState)
  let status;
  if (winner) {
    status = "Winner: " + winner;
  }

  return (
    <>
      <div className="status">{status}</div>
      <div className="boardRow">
        <Square value={boardState[0]} onClick={handleClick(0)}/>
        <Square value={boardState[1]} onClick={handleClick(1)}/>
        <Square value={boardState[2]} onClick={handleClick(2)}/>
      </div>
      <div className="boardRow">
        <Square value={boardState[3]} onClick={handleClick(3)}/>
        <Square value={boardState[4]} onClick={handleClick(4)}/>
        <Square value={boardState[5]} onClick={handleClick(5)}/>
      </div>
      <div className="boardRow">
        <Square value={boardState[6]} onClick={handleClick(6)}/>
        <Square value={boardState[7]} onClick={handleClick(7)}/>
        <Square value={boardState[8]} onClick={handleClick(8)}/>
      </div>
    </>
  )
}

function Square({value, onClick}) {
  return <button className="square" onClick={onClick}>{value}</button>;
}


function calculateWinner(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6]
  ];
  for (let i = 0; i < lines.length; i++) {
    const [a, b, c] = lines[i];
    if (squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
}
