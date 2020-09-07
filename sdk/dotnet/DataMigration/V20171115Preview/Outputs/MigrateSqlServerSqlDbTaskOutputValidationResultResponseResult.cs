// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20171115Preview.Outputs
{

    [OutputType]
    public sealed class MigrateSqlServerSqlDbTaskOutputValidationResultResponseResult
    {
        /// <summary>
        /// Result identifier
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Result type
        /// </summary>
        public readonly string ResultType;

        [OutputConstructor]
        private MigrateSqlServerSqlDbTaskOutputValidationResultResponseResult(
            string id,

            string resultType)
        {
            Id = id;
            ResultType = resultType;
        }
    }
}