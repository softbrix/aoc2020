
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

let USED = {};
let start = [];
rl.on('line', (line) => {
  start = line.split(',');
});

rl.on('close', (input) => {
  let i = 1;
  let last = 0;
  start.forEach(t => {
    USED[t] = i++;
  });
  for (; i < 30000000; ++i) {
    if (USED[last] == undefined) {
      USED[last] = i
      last = 0;
    } else {
      lNew = i - USED[last];
      USED[last] = i;
      last = lNew;
    }
    if ((i % 1000000) == 0) {
      process.stdout.write('.');
    }
    if (i == 2019) {
      console.log("Part 1: ", last)
    }
  }
  console.log("\nPart 2: ", last)
});
