import React from 'react'

interface Props {
  msg: string
}

export default function DefaultError({ msg }: Props) {
  return msg
}
