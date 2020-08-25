// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.V20190901.Outputs
{

    [OutputType]
    public sealed class ContactDetailsResponseResult
    {
        /// <summary>
        /// Contact name of the person.
        /// </summary>
        public readonly string ContactName;
        /// <summary>
        /// List of Email-ids to be notified about job progress.
        /// </summary>
        public readonly ImmutableArray<string> EmailList;
        /// <summary>
        /// Mobile number of the contact person.
        /// </summary>
        public readonly string? Mobile;
        /// <summary>
        /// Notification preference for a job stage.
        /// </summary>
        public readonly ImmutableArray<Outputs.NotificationPreferenceResponseResult> NotificationPreference;
        /// <summary>
        /// Phone number of the contact person.
        /// </summary>
        public readonly string Phone;
        /// <summary>
        /// Phone extension number of the contact person.
        /// </summary>
        public readonly string? PhoneExtension;

        [OutputConstructor]
        private ContactDetailsResponseResult(
            string contactName,

            ImmutableArray<string> emailList,

            string? mobile,

            ImmutableArray<Outputs.NotificationPreferenceResponseResult> notificationPreference,

            string phone,

            string? phoneExtension)
        {
            ContactName = contactName;
            EmailList = emailList;
            Mobile = mobile;
            NotificationPreference = notificationPreference;
            Phone = phone;
            PhoneExtension = phoneExtension;
        }
    }
}