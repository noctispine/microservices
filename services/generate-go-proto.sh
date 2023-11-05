protoBuffersDir="proto-buffers"
protoFileNames=`ls ${protoBuffersDir}`
baseProtoFile="base.proto"
baseProtoBufferDir="proto-buffers/baseProto"
baseProtoGatewayPath="./gateway/pkg"
for file in $protoFileNames
do
    fileName=`basename $file .proto`
    if [ $fileName == "base" ]; then
        continue
    fi
    
    serviceName=${fileName%.*}
    gatewayOutDir="./gateway/pkg/${serviceName}"
    serviceOutDir="./${serviceName}/pkg"

    protoc \
    --go_out=${gatewayOutDir} \
    --go_opt=M${file}.proto=github.com/capstone-project-bunker/backend/services/gateway/pkg/${serviceName}/pb/${filename} \
    --go-grpc_out=require_unimplemented_servers=false:${gatewayOutDir} \
    --go-grpc_opt=M${file}=github.com/capstone-project-bunker/backend/services/gateway/pkg/${serviceName}/pb \
    ${protoBuffersDir}/${file}
    
    protoc \
    --go_out=${serviceOutDir} \
    --go_opt=M${file}=github.com/capstone-project-bunker/backend/services/gateway/pkg/${serviceName}/pb/${filename} \
    --go-grpc_out=require_unimplemented_servers=false:${serviceOutDir} \
    --go-grpc_opt=M${file}=github.com/capstone-project-bunker/backend/services/${serviceName}/pkg/pb \
    ${protoBuffersDir}/${file}

    # protoc --go_out=${serviceOutDir} \
    # --go_opt=M${file}=github.com/capstone-project-bunker/backend/services/${serviceName}/baseProto \
    # --go-grpc_out=${serviceOutDir} --go-grpc_opt=module=github.com/capstone-project-bunker/backend/services/${serviceName} \
    # ${baseProtoBufferDir}/${baseProtoFile}
done

# protoc --go_out=${baseProtoGatewayPath} \
#     --go_opt=M${file}=github.com/capstone-project-bunker/backend/services/gateway/pkg/baseProto \
#     --go-grpc_out=${baseProtoGatewayPath} --go-grpc_opt=module=github.com/capstone-project-bunker/backend/services/gateway/pkg/baseProto \
#     ${baseProtoBufferDir}/${baseProtoFile}

exit 0 