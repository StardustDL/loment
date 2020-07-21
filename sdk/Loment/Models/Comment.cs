using System;
using System.Collections.Generic;
using System.Text;

namespace Loment.Models
{
    public class Comment
    {
        public string Id { get; set; } = string.Empty;

        public DateTimeOffset CreationTime { get; set; }

        public DateTimeOffset ModificationTime { get; set; }

        public string Content { get; set; } = string.Empty;

        public string Uri { get; set; } = string.Empty;

        public string Author { get; set; } = string.Empty;

        public string Email { get; set; } = string.Empty;

        public string Link { get; set; } = string.Empty;

        public string Extra { get; set; } = string.Empty;
    }
}
