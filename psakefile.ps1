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
    $env:CGO_ENABLED = 0
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
    Set-Location ./dist
    $env:LOMENT_DBORIGIN = "root:123456@(localhost:3306)"
    $env:LOMENT_DBNAME = "loment_db"
    $env:LOMENT_PORT = "4000"
    Exec { ./loment.exe }
}

Task Docker {
    Exec { docker-compose up -d db }
    Start-Sleep -Seconds 10
    Exec { docker-compose up -d }
    Start-Sleep -Seconds 5
}