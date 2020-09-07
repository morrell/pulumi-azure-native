// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The Edifact agreement protocol settings.
    /// </summary>
    public sealed class EdifactProtocolSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The EDIFACT acknowledgement settings.
        /// </summary>
        [Input("acknowledgementSettings", required: true)]
        public Input<Inputs.EdifactAcknowledgementSettingsArgs> AcknowledgementSettings { get; set; } = null!;

        [Input("edifactDelimiterOverrides")]
        private InputList<Inputs.EdifactDelimiterOverrideArgs>? _edifactDelimiterOverrides;

        /// <summary>
        /// The EDIFACT delimiter override settings.
        /// </summary>
        public InputList<Inputs.EdifactDelimiterOverrideArgs> EdifactDelimiterOverrides
        {
            get => _edifactDelimiterOverrides ?? (_edifactDelimiterOverrides = new InputList<Inputs.EdifactDelimiterOverrideArgs>());
            set => _edifactDelimiterOverrides = value;
        }

        [Input("envelopeOverrides")]
        private InputList<Inputs.EdifactEnvelopeOverrideArgs>? _envelopeOverrides;

        /// <summary>
        /// The EDIFACT envelope override settings.
        /// </summary>
        public InputList<Inputs.EdifactEnvelopeOverrideArgs> EnvelopeOverrides
        {
            get => _envelopeOverrides ?? (_envelopeOverrides = new InputList<Inputs.EdifactEnvelopeOverrideArgs>());
            set => _envelopeOverrides = value;
        }

        /// <summary>
        /// The EDIFACT envelope settings.
        /// </summary>
        [Input("envelopeSettings", required: true)]
        public Input<Inputs.EdifactEnvelopeSettingsArgs> EnvelopeSettings { get; set; } = null!;

        /// <summary>
        /// The EDIFACT framing settings.
        /// </summary>
        [Input("framingSettings", required: true)]
        public Input<Inputs.EdifactFramingSettingsArgs> FramingSettings { get; set; } = null!;

        /// <summary>
        /// The EDIFACT message filter.
        /// </summary>
        [Input("messageFilter", required: true)]
        public Input<Inputs.EdifactMessageFilterArgs> MessageFilter { get; set; } = null!;

        [Input("messageFilterList")]
        private InputList<Inputs.EdifactMessageIdentifierArgs>? _messageFilterList;

        /// <summary>
        /// The EDIFACT message filter list.
        /// </summary>
        public InputList<Inputs.EdifactMessageIdentifierArgs> MessageFilterList
        {
            get => _messageFilterList ?? (_messageFilterList = new InputList<Inputs.EdifactMessageIdentifierArgs>());
            set => _messageFilterList = value;
        }

        /// <summary>
        /// The EDIFACT processing Settings.
        /// </summary>
        [Input("processingSettings", required: true)]
        public Input<Inputs.EdifactProcessingSettingsArgs> ProcessingSettings { get; set; } = null!;

        [Input("schemaReferences", required: true)]
        private InputList<Inputs.EdifactSchemaReferenceArgs>? _schemaReferences;

        /// <summary>
        /// The EDIFACT schema references.
        /// </summary>
        public InputList<Inputs.EdifactSchemaReferenceArgs> SchemaReferences
        {
            get => _schemaReferences ?? (_schemaReferences = new InputList<Inputs.EdifactSchemaReferenceArgs>());
            set => _schemaReferences = value;
        }

        [Input("validationOverrides")]
        private InputList<Inputs.EdifactValidationOverrideArgs>? _validationOverrides;

        /// <summary>
        /// The EDIFACT validation override settings.
        /// </summary>
        public InputList<Inputs.EdifactValidationOverrideArgs> ValidationOverrides
        {
            get => _validationOverrides ?? (_validationOverrides = new InputList<Inputs.EdifactValidationOverrideArgs>());
            set => _validationOverrides = value;
        }

        /// <summary>
        /// The EDIFACT validation settings.
        /// </summary>
        [Input("validationSettings", required: true)]
        public Input<Inputs.EdifactValidationSettingsArgs> ValidationSettings { get; set; } = null!;

        public EdifactProtocolSettingsArgs()
        {
        }
    }
}