Task default -depends Restore, Build

Task Deploy {
    
}

Task CI -depends Install-deps, Restore, Build, Test, Benchmark, Report

Task CD -depends CD-Build, Deploy

Task CD-Build -depends Install-deps, Restore, Build

Task Restore {
}

Task Build {
    if (-not (Test-Path -Path "dist")) {
        New-Item -Path "dist" -ItemType Directory
    }
    Set-Location src
    Exec { go build -v -o ../dist/loment.exe }
    Set-Location ..
}

Task Install-deps {
}

Task Test {
    Set-Location src
    go test -v .
    Set-Location ..
}

Task Benchmark { 
    
}

Task Report {
}

Task Run -depends Build {
    # docker run --rm --security-opt seccomp:unconfined  -p 3306:3306 --name mymysql -e MYSQL_ROOT_PASSWORD=123456 mysql:latest
    Set-Location ./dist
    $env:LOMENT_DBORIGIN = "root:123456@(localhost:3306)"
    $env:LOMENT_DBNAME = "loment_db"
    $env:LOMENT_PORT = "4000"
    Exec { ./loment }
}