awk '$1 == "CURRENT_BUILD" {$3=($3+1)";"}1' projectbuild > projectbuild.new
mv projectbuild.new projectbuild
cat projectbuild | sed 's/;//g' |  awk '{print $3}' 
