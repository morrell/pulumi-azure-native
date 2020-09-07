// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180715Preview
{
    /// <summary>
    /// A file resource
    /// </summary>
    public partial class File : Pulumi.CustomResource
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom file properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ProjectFilePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a File resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public File(string name, FileArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datamigration/v20180715preview:File", name, args ?? new FileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private File(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datamigration/v20180715preview:File", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing File resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static File Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new File(name, id, options);
        }
    }

    public sealed class FileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Name of the File
        /// </summary>
        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Name of the project
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Custom file properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ProjectFilePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public FileArgs()
        {
        }
    }
}