# loment

![CI](https://github.com/StardustDL/loment/workflows/CI/badge.svg) ![CD](https://github.com/StardustDL/loment/workflows/CD/badge.svg) ![License](https://img.shields.io/github/license/StardustDL/loment.svg)

Loment is a Light cOmMENT service.

## API

- **Post** `/` with *Comment* body: Create comment, return id
- **Post** `/query` with *CommentQuery* body: Query comments, return list of comments
- **Get** `/id`: Get comment by id
- **Delete** `/id`: Delete comment by id
- **Put** `/id` with *Comment* body: Update comment by id

## Models

```go
type Comment struct {
	ID               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Location         string
	Nickname         string
	Mail             string
	Link             string
}

type CommentQuery struct {
	ID               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Location         string
	Nickname         string
	Mail             string
	Link             string
	Offset           int
	Limit            int
}
```

## SDK

For C#.

```sh
dotnet add package Loment
```

API:

```csharp
public interface ILomentService
{
    Task<string?> Create(Comment comment, CancellationToken cancellationToken = default);

    Task<IList<Comment>> Query(CommentQuery query, CancellationToken cancellationToken = default);

    Task<Comment?> Get(string id, CancellationToken cancellationToken = default);

    Task<bool> Delete(string id, CancellationToken cancellationToken = default);

    Task<bool> Update(Comment comment, CancellationToken cancellationToken = default);
}
```