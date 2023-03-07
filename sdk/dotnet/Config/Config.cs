// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Splunk
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("splunk");

        private static readonly __Value<string?> _authToken = new __Value<string?>(() => __config.Get("authToken"));
        /// <summary>
        /// Authentication tokens, also known as JSON Web Tokens (JWT), are a method for authenticating Splunk platform users into
        /// the Splunk platform
        /// </summary>
        public static string? AuthToken
        {
            get => _authToken.Get();
            set => _authToken.Set(value);
        }

        private static readonly __Value<bool?> _insecureSkipVerify = new __Value<bool?>(() => __config.GetBoolean("insecureSkipVerify"));
        /// <summary>
        /// insecure skip verification flag
        /// </summary>
        public static bool? InsecureSkipVerify
        {
            get => _insecureSkipVerify.Get();
            set => _insecureSkipVerify.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password"));
        /// <summary>
        /// Splunk instance password
        /// </summary>
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<int?> _timeout = new __Value<int?>(() => __config.GetInt32("timeout"));
        /// <summary>
        /// Timeout when making calls to Splunk server. Defaults to 60 seconds
        /// </summary>
        public static int? Timeout
        {
            get => _timeout.Get();
            set => _timeout.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url"));
        /// <summary>
        /// Splunk instance URL
        /// </summary>
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

        private static readonly __Value<string?> _username = new __Value<string?>(() => __config.Get("username"));
        /// <summary>
        /// Splunk instance admin username
        /// </summary>
        public static string? Username
        {
            get => _username.Get();
            set => _username.Set(value);
        }

    }
}
