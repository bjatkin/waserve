import { useEffect, useState, useRef } from 'react'
import Grid from './components/grid'
import './App.css'

function App() {
  const tileSize:number = 110;
  
  const [size, setSize] = useState({
    height: Math.floor(window.innerHeight / tileSize),
    width: Math.floor(window.innerWidth / tileSize)
  })

  useEffect(() => {
    const handleResize = () => {
      setSize({
        height: Math.floor(window.innerHeight / tileSize),
        width: Math.floor(window.innerWidth / tileSize)
      })
    }

    window.addEventListener('resize', handleResize)

    return () => {
      window.removeEventListener('resize', handleResize)
    }
  }, [])

  const appRef = useRef<HTMLDivElement>(null)

  const handleClick = () => {
    if (appRef.current != null) {
      appRef.current.style.setProperty("--bg-1", "#833AB4")
      appRef.current.style.setProperty("--bg-2", "#FD1D1D")
      appRef.current.style.setProperty("--bg-3", "#FCB045")
      appRef.current.style.setProperty("--speed", "6s")
    }
  }

  return (
    <div className="App" ref={appRef}>
      <Grid
        columns={size.height}
        rows={size.width}
      ></Grid>
      <button id="runButton" onClick={handleClick}>
        <svg id="runArrow" height="40" width="80" viewBox='0 0 100 100'>
          <polygon id="runPolygon" points="0,0 140,50 0,100"></polygon>
        </svg>
      </button>
    </div>
  )
}

export default App
