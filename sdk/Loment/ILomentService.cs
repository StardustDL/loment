using Loment.Models;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Threading;
using System.Threading.Tasks;

namespace Loment
{
    public interface ILomentService
    {
        Task<string?> Create(Comment comment, CancellationToken cancellationToken = default);

        Task<IList<Comment>> Query(CommentQuery query, CancellationToken cancellationToken = default);

        Task<long> Count(CommentQuery query, CancellationToken cancellationToken = default);

        Task<Comment?> Get(string id, CancellationToken cancellationToken = default);

        Task<bool> Delete(string id, CancellationToken cancellationToken = default);

        Task<bool> Update(Comment comment, CancellationToken cancellationToken = default);
    }

    public class LomentService : ILomentService
    {
        public LomentService(HttpClient client) => Client = client;

        public HttpClient Client { get; }

        public async Task<string?> Create(Comment comment, CancellationToken cancellationToken = default)
        {
            var response = await Client.PostAsJsonAsync("/", comment, cancellationToken).ConfigureAwait(false);
            response.EnsureSuccessStatusCode();
            return await response.Content.ReadAsStringAsync();
        }

        public async Task<bool> Delete(string id, CancellationToken cancellationToken = default)
        {
            var response = await Client.DeleteAsync($"/{Uri.EscapeDataString(id)}", cancellationToken).ConfigureAwait(false);
            response.EnsureSuccessStatusCode();
            return await response.Content.ReadFromJsonAsync<bool>(cancellationToken: cancellationToken);
        }

        public async Task<Comment?> Get(string id, CancellationToken cancellationToken = default)
        {
            var response = await Client.GetAsync($"/{Uri.EscapeDataString(id)}", cancellationToken).ConfigureAwait(false);
            response.EnsureSuccessStatusCode();
            return await response.Content.ReadFromJsonAsync<Comment?>(cancellationToken: cancellationToken);
        }

        public async Task<IList<Comment>> Query(CommentQuery query, CancellationToken cancellationToken = default)
        {
            var response = await Client.PostAsJsonAsync("/query", query, cancellationToken).ConfigureAwait(false);
            response.EnsureSuccessStatusCode();
            var result = await response.Content.ReadFromJsonAsync<IList<Comment>>(cancellationToken: cancellationToken);
            if(result is null)
            {
                return Array.Empty<Comment>();
            }
            else
            {
                return result;
            }
        }

        public async Task<long> Count(CommentQuery query, CancellationToken cancellationToken = default)
        {
            var response = await Client.PostAsJsonAsync("/count", query, cancellationToken).ConfigureAwait(false);
            response.EnsureSuccessStatusCode();
            var result = await response.Content.ReadFromJsonAsync<long>(cancellationToken: cancellationToken);
            return result;
        }

        public async Task<bool> Update(Comment comment, CancellationToken cancellationToken = default)
        {
            var response = await Client.PutAsJsonAsync($"/{Uri.EscapeDataString(comment.Id)}", comment, cancellationToken).ConfigureAwait(false);
            response.EnsureSuccessStatusCode();
            return await response.Content.ReadFromJsonAsync<bool>(cancellationToken: cancellationToken);
        }
    }
}
