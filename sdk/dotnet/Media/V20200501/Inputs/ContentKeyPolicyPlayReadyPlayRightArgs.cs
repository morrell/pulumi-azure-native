// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20200501.Inputs
{

    /// <summary>
    /// Configures the Play Right in the PlayReady license.
    /// </summary>
    public sealed class ContentKeyPolicyPlayReadyPlayRightArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures Automatic Gain Control (AGC) and Color Stripe in the license. Must be between 0 and 3 inclusive.
        /// </summary>
        [Input("agcAndColorStripeRestriction")]
        public Input<int>? AgcAndColorStripeRestriction { get; set; }

        /// <summary>
        /// Configures Unknown output handling settings of the license.
        /// </summary>
        [Input("allowPassingVideoContentToUnknownOutput", required: true)]
        public Input<string> AllowPassingVideoContentToUnknownOutput { get; set; } = null!;

        /// <summary>
        /// Specifies the output protection level for compressed digital audio.
        /// </summary>
        [Input("analogVideoOpl")]
        public Input<int>? AnalogVideoOpl { get; set; }

        /// <summary>
        /// Specifies the output protection level for compressed digital audio.
        /// </summary>
        [Input("compressedDigitalAudioOpl")]
        public Input<int>? CompressedDigitalAudioOpl { get; set; }

        /// <summary>
        /// Specifies the output protection level for compressed digital video.
        /// </summary>
        [Input("compressedDigitalVideoOpl")]
        public Input<int>? CompressedDigitalVideoOpl { get; set; }

        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        [Input("digitalVideoOnlyContentRestriction", required: true)]
        public Input<bool> DigitalVideoOnlyContentRestriction { get; set; } = null!;

        /// <summary>
        /// Configures the Explicit Analog Television Output Restriction in the license. Configuration data must be between 0 and 3 inclusive.
        /// </summary>
        [Input("explicitAnalogTelevisionOutputRestriction")]
        public Input<Inputs.ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionArgs>? ExplicitAnalogTelevisionOutputRestriction { get; set; }

        /// <summary>
        /// The amount of time that the license is valid after the license is first used to play content.
        /// </summary>
        [Input("firstPlayExpiration")]
        public Input<string>? FirstPlayExpiration { get; set; }

        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        [Input("imageConstraintForAnalogComponentVideoRestriction", required: true)]
        public Input<bool> ImageConstraintForAnalogComponentVideoRestriction { get; set; } = null!;

        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        [Input("imageConstraintForAnalogComputerMonitorRestriction", required: true)]
        public Input<bool> ImageConstraintForAnalogComputerMonitorRestriction { get; set; } = null!;

        /// <summary>
        /// Configures the Serial Copy Management System (SCMS) in the license. Must be between 0 and 3 inclusive.
        /// </summary>
        [Input("scmsRestriction")]
        public Input<int>? ScmsRestriction { get; set; }

        /// <summary>
        /// Specifies the output protection level for uncompressed digital audio.
        /// </summary>
        [Input("uncompressedDigitalAudioOpl")]
        public Input<int>? UncompressedDigitalAudioOpl { get; set; }

        /// <summary>
        /// Specifies the output protection level for uncompressed digital video.
        /// </summary>
        [Input("uncompressedDigitalVideoOpl")]
        public Input<int>? UncompressedDigitalVideoOpl { get; set; }

        public ContentKeyPolicyPlayReadyPlayRightArgs()
        {
        }
    }
}