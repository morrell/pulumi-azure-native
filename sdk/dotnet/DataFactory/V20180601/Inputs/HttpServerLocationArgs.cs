// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20180601.Inputs
{

    /// <summary>
    /// The location of http server.
    /// </summary>
    public sealed class HttpServerLocationArgs : Pulumi.ResourceArgs
    {
        [Input("fileName")]
        private InputMap<object>? _fileName;

        /// <summary>
        /// Specify the file name of dataset. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> FileName
        {
            get => _fileName ?? (_fileName = new InputMap<object>());
            set => _fileName = value;
        }

        [Input("folderPath")]
        private InputMap<object>? _folderPath;

        /// <summary>
        /// Specify the folder path of dataset. Type: string (or Expression with resultType string)
        /// </summary>
        public InputMap<object> FolderPath
        {
            get => _folderPath ?? (_folderPath = new InputMap<object>());
            set => _folderPath = value;
        }

        [Input("relativeUrl")]
        private InputMap<object>? _relativeUrl;

        /// <summary>
        /// Specify the relativeUrl of http server. Type: string (or Expression with resultType string)
        /// </summary>
        public InputMap<object> RelativeUrl
        {
            get => _relativeUrl ?? (_relativeUrl = new InputMap<object>());
            set => _relativeUrl = value;
        }

        /// <summary>
        /// Type of dataset storage location.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public HttpServerLocationArgs()
        {
        }
    }
}