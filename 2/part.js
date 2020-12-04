
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

let sum1 = 0;
let sum2 = 0;
rl.on('line', (line) => {
  l = line.split(' ');
  [lo, hi] = l[0].split('-');
  c = l[1].substr(0, 1);
  chars = l[2];
  tot = chars.split('').reduce((t, s) => t + (s === c ? 1 : 0), 0)
  if (lo <= tot && tot <= hi) {
    sum1 += 1;
  }
  if (chars[lo-1] === c ^ chars[hi-1] === c) {
    sum2 += 1;
  }
});

rl.on('close', (input) => {
  console.log(`Sum1: ${sum1}`);
  console.log(`Sum2: ${sum2}`);
});
