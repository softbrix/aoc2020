
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

const MY_BAG = "shiny gold"

let sum1 = 0;
let sum2 = 0;
let BAGS = {}
rl.on('line', (line) => {
  l = line.split(' bags contain ');
  id = l[0]
  BAGS[id] = l[1].split(', ').map(t => {q = t.split(' '); return {"c": parseInt(q[0]), b: q[1] + ' ' + q[2]}})
});

rl.on('close', (input) => {
  let visited = {};
  let bfs_visited = {};
  sum1 = Object.keys(BAGS).filter(b => b != MY_BAG).reduce((a,b) => a+red(b, visited), 0);
  sum2 = bfs(MY_BAG, bfs_visited);
  console.log(`Sum1: ${sum1} of ${Object.keys(BAGS).length}`);
  console.log(`Sum2: ${sum2}`);
});

function red(b, visited) {
  if (isNaN(BAGS[b][0].c)) {
    return 0;
  }
  if (b === MY_BAG) {
    return 1;
  }
  if (visited[b] === undefined) {
    visited[b] = BAGS[b].reduce((a,b) => a + red(b.b, visited), 0);
  } 
  return visited[b] > 0 ? 1 : 0; 
}

function bfs(bag, visited) { 
  if (isNaN(BAGS[bag][0].c)) {
    return 0;
  }
  if (visited[bag] === undefined) {
    visited[bag] = BAGS[bag].reduce((a,b) => a + b.c + b.c * bfs(b.b, visited), 0);
  } 
  return visited[bag]; 
} 
