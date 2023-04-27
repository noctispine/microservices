const { Command } = require("commander")
const figlet = require("figlet")
const { create } = require("./commands/create")

console.log(figlet.textSync("capstone - mc"))

const program = new Command();

program
  .name('capstone-service')
  .description('CLI to init & startup a microservice in monorepo')
  .description('Split a string into substrings and display as an array')

program
  .command('create')
  .argument('<projectName>')
  .argument('<port>')
  .action(() => {
    create(program)
  })

program.parse();





