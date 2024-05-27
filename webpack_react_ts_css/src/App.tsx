import { useState } from "react"
import { reverseStr } from "./someLib"

export default function Form() {
  const startingStr = "Change me"

  const [inValue, setInValue] = useState(startingStr)

  const [outValue, setOutValue] = useState(
    reverseStr(startingStr)
  )

  function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
    setInValue(event.currentTarget.value)
    setOutValue(
      reverseStr(event.currentTarget.value)
    )
  }

  return (
    <>
      <input value={inValue} onChange={handleChange} />
      <p>Value: {outValue}</p>
    </>
  );
}
