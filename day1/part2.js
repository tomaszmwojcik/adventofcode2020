// Now we need to find three numbers that sums up to something else

const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

let entries = [];

rl.on('line', (input) => {
  const newEntry = parseInt(input); 
  for(let i=0; i<entries.length; i++) {
      for(let j=i+1; j<entries.length; j++) {
        if(newEntry + entries[i] + entries[j] === 2020) {
            console.log(`Found result: ${newEntry * entries[i] * entries[j]}`);
            rl.close();
            return;
        }
      }
  }
  entries.push(newEntry);
});