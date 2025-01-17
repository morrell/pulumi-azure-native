#!/bin/bash
set -euf -o pipefail

currDateTime=$(date +"%Y-%m-%d_%H-%M-%S")

# Makes sure that each summary has a unique name in S3
newSummaryName="summary_${currDateTime}.json"
echo "New file to be uploaded: ${newSummaryName}"

# Sets up URI of S3 bucket and location to upload to
# Example URI: s3://arm2pulumi-coverage-results-c9610a2/summaries/summary_2021-05-11_19-44-12.json
s3BucketName="arm2pulumi-coverage-results-d127fff"
s3KeyName="summaries/${newSummaryName}"
s3FullURI="s3://${s3BucketName}/${s3KeyName}"

cd test-results
# Edit JSON summary file to be copiable into Redshift
# Changing file from a list of JSON objects to group of JSON objects. Done by 
# removing first ("[") and last line ("]") and replacing all "}," with "}".
# Modified JSON contents are then copied into the file `newSummaryName`.
sed '1d' summary.json | sed '$d' | sed -E 's/},/}/' > $newSummaryName

aws s3 cp $newSummaryName $s3FullURI --acl bucket-owner-full-control
