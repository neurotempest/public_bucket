/**
 * @jest-environment jsdom
 */
import { render, screen } from '@testing-library/react';


import Form from "../src/App"

describe("App", () => {
  it("does something", () => {
    render(<Form />)
    screen.debug()
  })
})
