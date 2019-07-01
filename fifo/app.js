var readline = require('readline');
var rl = readline.createInterface({
  input: process.stdin,
  terminal: false
});

rl.on('line', function(line){
    // Send log message to stderr
    process.stderr.write("Input received");
    // Send output to stdout
    process.stdout.write("NODE! " + line + "\n");
})
