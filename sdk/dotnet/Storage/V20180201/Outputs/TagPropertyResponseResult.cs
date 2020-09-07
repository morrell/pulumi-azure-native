// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20180201.Outputs
{

    [OutputType]
    public sealed class TagPropertyResponseResult
    {
        /// <summary>
        /// Returns the Object ID of the user who added the tag.
        /// </summary>
        public readonly string ObjectIdentifier;
        /// <summary>
        /// The tag value.
        /// </summary>
        public readonly string Tag;
        /// <summary>
        /// Returns the Tenant ID that issued the token for the user who added the tag.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// Returns the date and time the tag was added.
        /// </summary>
        public readonly string Timestamp;
        /// <summary>
        /// Returns the User Principal Name of the user who added the tag.
        /// </summary>
        public readonly string Upn;

        [OutputConstructor]
        private TagPropertyResponseResult(
            string objectIdentifier,

            string tag,

            string tenantId,

            string timestamp,

            string upn)
        {
            ObjectIdentifier = objectIdentifier;
            Tag = tag;
            TenantId = tenantId;
            Timestamp = timestamp;
            Upn = upn;
        }
    }
}