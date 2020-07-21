using System;

namespace Loment.Models
{
    public class CommentQuery
    {
        public string? ID { get; set; }

        public DateTimeOffset? CreationTime { get; set; }

        public DateTimeOffset? ModificationTime { get; set; }

        public string? Content { get; set; }

        public string? Uri { get; set; }

        public string? Author { get; set; }

        public string? Email { get; set; }

        public string? Link { get; set; }

        public int Offset { get; set; } = 0;

        public int Limit { get; set; } = 10;
    }
}
