import * as _ from "lodash"

export function reverseStr(
  str: string,
): string {

  return _.reverse(str.split("")).join('')
}
