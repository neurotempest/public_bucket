import {
  reverseStr
} from "../src/someLib"

describe("reverseStr", () => {

  it("reverses a string", () => {
    expect(reverseStr("abcde fgh")).toEqual("hgf edcba")
  })
  it("reverse another string", () => {
    expect(reverseStr("some string")).toEqual("gnirts emos")
  })
})
