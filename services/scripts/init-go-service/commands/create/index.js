const fs = require("fs")
const path = require("path")
const chalk = require("chalk")
const { getAllFiles } = require("./helper")

async function create(program) {
    const projectName = program.args[1]
    const port = program.args[2]
    copyProject(`./${projectName}`)
    const files = getAllFiles(`./${projectName}`)
    console.log(files)

    files.forEach(file => {
        fs.readFile(file, 'utf-8', (err, data) => {
            if(err) {
                console.error("error occured while writing to file", err)
                process.exit(1)
            }

            const result = data
                .replace(/\$serviceNameCapitalized/g, `${projectName.charAt(0).toUpperCase() + projectName.slice(1)}`)
                .replace(/\$serviceName/g, projectName)
                .replace(/\$PORT/g, port)
            
            fs.writeFile(file, result, 'utf-8', err => {
                if(err) {
                    console.error("error occured while writing to file", err)
                    process.exit(1)
                }
            })

        })


    })
    console.log(chalk.yellow(projectName) + chalk.cyan(" service is created"))

}

function copyProject(dest) {
    fs.cpSync('./scripts/init-go-service/service', dest, {recursive: true}, (err) => {
        console.error("service cannot copied to destination")
        process.exitCode(2)
    })
}

module.exports = { create }