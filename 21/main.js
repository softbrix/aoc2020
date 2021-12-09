
const readline = require('readline')

const rl = readline.createInterface({
  input: process.stdin
});

let sum1 = 0;
let sum2 = 0;
let ingredients = {}
let matches = {}
rl.on('line', (line) => {
  if (line.length == 0) {
    return
  }
  line = line.slice(0, -1)
  let [ing, allergenes] = line.split(' (contains ')
  ing = ing.split(' ')
  allergenes = allergenes.split(', ')
  // console.log(ing, allergenes)
  ing.forEach(ing => {
    ingredients[ing] = ingredients[ing] || 0
    ingredients[ing]++
  })
  allergenes.forEach(a => {
    if (matches[a] !== undefined) {
      Object.keys(matches[a]).forEach(k => {
        if (ing.indexOf(k) < 0) {
          delete matches[a][k]
        }
      });
    } else {
      matches[a] = {}
      ing.forEach(i => {
        matches[a][i] = true
      })
    }
  })
});

rl.on('close', (input) => {
  const MAX_LOOP = 20;
  let i = 0
  let found = []
  let allergenMatch = {}
  while (Object.keys(matches).length > 0 && i++ < MAX_LOOP) {
    Object.entries(matches).forEach(ent => {
      let [key, values] = ent
      let ingKeys = Object.keys(values)
      if (ingKeys.length == 1) {
        let ing = ingKeys[0]
        if (found.indexOf(ing) >= 0) {
          console.error("Already found: ", ing)
          process.exit(1)
        }
        found.push(ing)
        allergenMatch[key] = ing
        delete matches[key]
      } else {
        found.forEach(ing => {
          delete matches[key][ing]
        })
      }
    })
  }

  if (i == MAX_LOOP) {
    console.error("Max loop")
    process.exit(1)
  }

  sum1 = Object.entries(ingredients).reduce((acc, [key, value]) => {
    if (found.indexOf(key) < 0) {
      acc += value
    }
    return acc
  }, 0)

  sum2 = Object.keys(allergenMatch).sort().reduce((acc, key) => {
    acc += allergenMatch[key] + ","
    return acc
  },  "").slice(0, -1)

  console.log(`Sum1: ${sum1}`);
  console.log(`Sum2: ${sum2}`);
});
