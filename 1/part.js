
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

let sum = 0
let rows = [];
rl.on('line', (line) => {
  n = Number(line);
  for (let x in rows) {
    for (let y in rows) {
      if (rows[y]+rows[x]+n === 2020) {
        console.log('result', rows[y]*rows[x]*line);
      }
    }
  }
  rows.push(n);
});

rl.on('close', (input) => {
  console.log(`Sum: ${sum}`);
});
