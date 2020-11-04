# Build it
go build -o terraform-provider-stitch_v0.1.0
if [ $? != 0 ]; then
  echo Error compiling. Return code: $?
  exit
fi
mv terraform-provider-stitch_v0.1.0 ~/.terraform.d/plugins/stitchdata.com/provider/stitch/0.1.0/darwin_amd64/terraform-provider-stitch_v0.1.0
rm -Rf ~/.terraform.d/plugin-cache/stitchdata.com

# Run it
cd terraform
rm -Rf .terraform terraform.tfstate terraform.tfstate.backup
terraform init
terraform plan
terraform apply --auto-approve
cd ..
