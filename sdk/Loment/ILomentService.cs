using Loment.Models;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

namespace Loment
{
    public interface ILomentService
    {
        Task<string?> Create(Comment comment, CancellationToken cancellationToken = default);

        Task<IList<Comment>> Query(CommentQuery query, CancellationToken cancellationToken = default);

        Task<Comment?> Get(string id, CancellationToken cancellationToken = default);

        Task<bool> Delete(string id, CancellationToken cancellationToken = default);

        Task<bool> Update(Comment comment, CancellationToken cancellationToken = default);
    }
}
