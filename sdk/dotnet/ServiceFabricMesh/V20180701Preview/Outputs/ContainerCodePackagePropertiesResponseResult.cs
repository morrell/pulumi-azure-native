// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class ContainerCodePackagePropertiesResponseResult
    {
        /// <summary>
        /// Command array to execute within the container in exec form.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// Reference to sinks in DiagnosticsDescription.
        /// </summary>
        public readonly Outputs.DiagnosticsRefResponseResult? Diagnostics;
        /// <summary>
        /// The endpoints exposed by this container.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointPropertiesResponseResult> Endpoints;
        /// <summary>
        /// Override for the default entry point in the container.
        /// </summary>
        public readonly string? Entrypoint;
        /// <summary>
        /// The environment variables to set in this container
        /// </summary>
        public readonly ImmutableArray<Outputs.EnvironmentVariableResponseResult> EnvironmentVariables;
        /// <summary>
        /// The Container image to use.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// Image registry credential.
        /// </summary>
        public readonly Outputs.ImageRegistryCredentialResponseResult? ImageRegistryCredential;
        /// <summary>
        /// Runtime information of a container instance.
        /// </summary>
        public readonly Outputs.ContainerInstanceViewResponseResult InstanceView;
        /// <summary>
        /// The labels to set in this container.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerLabelResponseResult> Labels;
        /// <summary>
        /// The name of the code package.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// This type describes the resource requirements for a container or a service.
        /// </summary>
        public readonly Outputs.ResourceRequirementsResponseResult Resources;
        /// <summary>
        /// The settings to set in this container. The setting file path can be fetched from environment variable "Fabric_SettingPath". The path for Windows container is "C:\\secrets". The path for Linux container is "/var/secrets".
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingResponseResult> Settings;
        /// <summary>
        /// The volumes to be attached to the container.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerVolumeResponseResult> VolumeRefs;

        [OutputConstructor]
        private ContainerCodePackagePropertiesResponseResult(
            ImmutableArray<string> commands,

            Outputs.DiagnosticsRefResponseResult? diagnostics,

            ImmutableArray<Outputs.EndpointPropertiesResponseResult> endpoints,

            string? entrypoint,

            ImmutableArray<Outputs.EnvironmentVariableResponseResult> environmentVariables,

            string image,

            Outputs.ImageRegistryCredentialResponseResult? imageRegistryCredential,

            Outputs.ContainerInstanceViewResponseResult instanceView,

            ImmutableArray<Outputs.ContainerLabelResponseResult> labels,

            string name,

            Outputs.ResourceRequirementsResponseResult resources,

            ImmutableArray<Outputs.SettingResponseResult> settings,

            ImmutableArray<Outputs.ContainerVolumeResponseResult> volumeRefs)
        {
            Commands = commands;
            Diagnostics = diagnostics;
            Endpoints = endpoints;
            Entrypoint = entrypoint;
            EnvironmentVariables = environmentVariables;
            Image = image;
            ImageRegistryCredential = imageRegistryCredential;
            InstanceView = instanceView;
            Labels = labels;
            Name = name;
            Resources = resources;
            Settings = settings;
            VolumeRefs = volumeRefs;
        }
    }
}