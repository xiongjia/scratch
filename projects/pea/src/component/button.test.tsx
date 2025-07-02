import { render } from '@testing-library/react'
import { describe, expect, it } from 'vitest'

import TestButton from './button'

describe('basic arithmetic checks', () => {
  it('1 + 1 equals 2', () => {
    expect(1 + 1).toBe(2)
  })

  it('2 * 2 equals 4', () => {
    render(<TestButton />)
  })
})
