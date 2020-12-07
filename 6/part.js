
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

function sumUp() {
  //console.log(cnt, answers, sum2)
  sum1 += Object.keys(answers).length;
  sum2 += Object.values(answers).reduce((a,t) => { a += (t == cnt ? 1 : 0); /*console.log(t, a)*/; return a }, 0)
  answers = {};
  cnt = 0;
}

let sum1 = 0
let sum2 = 0;
let answers = {};
let cnt = 0;
rl.on('line', (line) => {
  if (line.length == 0) {
    return sumUp();
  }
  cnt++;
  line.split('').forEach(t => { answers[t] = answers[t] || 0; answers[t]++});
});

rl.on('close', (input) => { 
  sumUp();
  console.log(`Sum1: ${sum1}`);
  console.log(`Sum2: ${sum2}`);
}); 
 