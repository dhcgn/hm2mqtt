New-Item .\build\package -ItemType Directory -ErrorAction Ignore
Copy-Item .\build\publish\GoHomeMaticMqtt_linux_arm .\packaging\bin

Remove-Item .\build\package\*

$version = Get-Content .\packaging\VERSION

tar -cvzf .\build\package\GoHomeMaticMqtt_$version.tar.gz -C packaging *