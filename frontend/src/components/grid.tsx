import { useEffect, useRef } from 'react'

interface GridProps {
  columns: number
  rows:    number 
}

function Grid(props: GridProps):JSX.Element {
  let tiles:JSX.Element[] = []
  for (let i = 0; i < props.rows*props.columns; i++) {
    tiles.push(<div className="tile" key={i.toString()}></div>)
  }

  const gridRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (gridRef.current != null) {
      gridRef.current.style.setProperty("--columns", props.columns.toString())
      gridRef.current.style.setProperty("--rows", props.rows.toString())
    }
  }, [])

  return (
    <div id="grid" ref={gridRef}>
      {tiles}
    </div>
  )
}

export default Grid