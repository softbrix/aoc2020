
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

function strToId(str) {
  let sum = 0;
  for (var i = 0; i < str.length; ++i) {
    if (str[i] == '#') {
      sum += Math.pow(2, i)
    }
  }
  return sum;
}

function reverse(str) {
  return str.split("").reverse().join("");
}

function column(arr, idx) {
  return arr.map(l => l[idx]).join('')
}

let sum1 = 0;
let sum2 = 0;
let tiles = {}
let active = 0;
rl.on('line', (line) => {
  if (line.length == 0) {
    active = -1
    return
  }
  if (line.startsWith('Tile ')) {
    active = line.substring('Tile '.length, line.length-1)
    return
  }
  tiles[active] = (tiles[active] || [])
  tiles[active].push(line)
});

let edges = {}
rl.on('close', (input) => {
  Object.keys(tiles).forEach(k => {
    tileEdges = [tiles[k][0], tiles[k][tiles[k].length-1], column(tiles[k],0), column(tiles[k],tiles[k][0].length-1)]
    tileEdges = tileEdges.reduce((acc, t) => {
      acc.push(t);
      acc.push(reverse(t))
      return acc
    }, [])
    tileEdges.forEach(e => {
      let id = strToId(e);
      edges[id] = edges[id] || []
      edges[id].push(k)
    })
  })

  edgeTiles = Object.entries(edges).reduce((acc, ent) => {
    let key = ent[0]
    let values = ent[1]
    if (values.length == 1) {
      acc[values[0]] = acc[values[0]] || []
      acc[values[0]].push(key)
    }
    return acc
  }, {})

  edgeTiles = Object.entries(edgeTiles).filter(f => {
    return f[1].length > 3
  })
  
  sum1 = edgeTiles.reduce((acc, t) => {
    return acc * t[0] 
  }, 1)

  console.log(edgeTiles)

  console.log(`Sum1: ${sum1}`);
  console.log(`Sum2: ${sum2}`);
});
