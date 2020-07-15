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
    go build -v -o ../dist/loment.exe
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
    Set-Location ./dist
    $env:LOMENT_PORT = "4000"
    Exec { ./loment }
}