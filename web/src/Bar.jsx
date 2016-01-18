/* @flow */
import React from 'react'

export default ({ percent }: { percent: number }): ReactElement => {
  return (
    <div style={BarStyles.container}>
      <span style={BarStyles.fill(percent)}>
        {percent}%
      </span>
    </div>
  )
}

let BarStyles = {
  container: {
    height: 20,
    position: 'relative',
    background: '#555',
    borderRadius: 5,
    padding: 5
  },
  fill: (percent: number) => ({
    display: 'block',
    width: percent + '%',
    height: '100%',
    borderRadius: 2.5,
    overflow: 'hidden',
    backgroundColor: '#31708F',
    color: '#FFF',
    display: 'flex',
    alignItems: 'center',
    paddingLeft: 5,
  })
}
