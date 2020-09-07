// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview.Outputs
{

    [OutputType]
    public sealed class ContentKeyPolicyFairPlayConfigurationResponseResult
    {
        /// <summary>
        /// The key that must be used as FairPlay ASk.
        /// </summary>
        public readonly string Ask;
        /// <summary>
        /// The Base64 representation of FairPlay certificate in PKCS 12 (pfx) format (including private key).
        /// </summary>
        public readonly string FairPlayPfx;
        /// <summary>
        /// The password encrypting FairPlay certificate in PKCS 12 (pfx) format.
        /// </summary>
        public readonly string FairPlayPfxPassword;
        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        public readonly string OdataType;
        /// <summary>
        /// The rental and lease key type.
        /// </summary>
        public readonly string RentalAndLeaseKeyType;
        /// <summary>
        /// The rental duration. Must be greater than or equal to 0.
        /// </summary>
        public readonly int RentalDuration;

        [OutputConstructor]
        private ContentKeyPolicyFairPlayConfigurationResponseResult(
            string ask,

            string fairPlayPfx,

            string fairPlayPfxPassword,

            string odataType,

            string rentalAndLeaseKeyType,

            int rentalDuration)
        {
            Ask = ask;
            FairPlayPfx = fairPlayPfx;
            FairPlayPfxPassword = fairPlayPfxPassword;
            OdataType = odataType;
            RentalAndLeaseKeyType = rentalAndLeaseKeyType;
            RentalDuration = rentalDuration;
        }
    }
}