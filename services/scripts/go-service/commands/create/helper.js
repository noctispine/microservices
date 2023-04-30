const fs = require("fs")
const path = require("path")

function getAllFiles(dirPath, arrayOfFiles) {
    files = fs.readdirSync(dirPath)

    arrayOfFiles = arrayOfFiles || []

    files.forEach(function(file) {
        if(fs.statSync(dirPath + "/" + file).isDirectory()) {
            arrayOfFiles = getAllFiles(dirPath + "/" + file, arrayOfFiles)
        } else {
            arrayOfFiles.push(path.join(dirPath, "/", file))
        }
    })

    return arrayOfFiles
}


function constructFileDirs(projectName) {
    return [
        {
            src: './scripts/go-service/files/service',
            dest: `./${projectName}`
        },
        {
            src: './scripts/go-service/files/example.proto',
            dest: `./proto-buffers/${projectName}.proto`
        }
    ]
}



module.exports = { getAllFiles, constructFileDirs }