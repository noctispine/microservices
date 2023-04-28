const { Command } = require("commander")
const figlet = require("figlet")
const { create } = require("./commands/create")
const chalk = require("chalk")

console.log(chalk.red(figlet.textSync("capstone - mc")))
console.log('\n')

const program = new Command();

program
  .name('capstone-service')
  .description('CLI to init & startup a microservice in monorepo')
  .description('Split a string into substrings and display as an array')

program
  .command('create')
  .description('create go microservice')
  .action(() => {
    create(program)
  })

program.parse();