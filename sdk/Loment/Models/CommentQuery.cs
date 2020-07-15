using System;

namespace Loment.Models
{
    public class CommentQuery
    {
        public string? ID { get; set; }

        public DateTimeOffset? CreationTime { get; set; }

        public DateTimeOffset? ModificationTime { get; set; }

        public string? Content { get; set; }

        public string? Location { get; set; }

        public string? Nickname { get; set; }

        public string? Mail { get; set; }

        public string? Link { get; set; }

        public int Offset { get; set; } = 0;

        public int Limit { get; set; } = 10;
    }
}
