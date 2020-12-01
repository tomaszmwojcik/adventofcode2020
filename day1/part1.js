// Finds two numbers that sum to some given number
// Think how this can be optimized(sorting the array?)

const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

let entries = [];

rl.on('line', (input) => {
  const newEntry = parseInt(input); 
  const matchingIndex = entries.findIndex(entry => newEntry + entry === 2020); 
  if(matchingIndex != -1) {
    console.log(`Found result: ${newEntry * entries[matchingIndex]}`);
    rl.close();
    return;
  }
  entries.push(newEntry);
});