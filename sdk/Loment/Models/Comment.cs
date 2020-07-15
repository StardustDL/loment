using System;
using System.Collections.Generic;
using System.Text;

namespace Loment.Models
{
    public class Comment
    {
        public string ID { get; set; } = string.Empty;

        public DateTimeOffset CreationTime { get; set; }

        public DateTimeOffset ModificationTime { get; set; }

        public string Content { get; set; } = string.Empty;

        public string Location { get; set; } = string.Empty;

        public string Nickname { get; set; } = string.Empty;

        public string Mail { get; set; } = string.Empty;

        public string Link { get; set; } = string.Empty;
    }
}
