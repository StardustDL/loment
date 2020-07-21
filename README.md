# loment

![CI](https://github.com/StardustDL/loment/workflows/CI/badge.svg) ![CD](https://github.com/StardustDL/loment/workflows/CD/badge.svg) ![License](https://img.shields.io/github/license/StardustDL/loment.svg) [![Loment](https://buildstats.info/nuget/Loment)](https://www.nuget.org/packages/Loment/)

Loment is a Light cOmMENT service.

## API

- **Post** `/` with *Comment* body: Create comment, return id
- **Post** `/query` with *CommentQuery* body: Query comments, return list of comments
- **Get** `/id`: Get comment by id, return comment
- **Delete** `/id`: Delete comment by id, return if done
- **Put** `/id` with *Comment* body: Update comment by id, return if done

## Models

```go
type Comment struct {
	Id               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Uri              string
	Author           string
	Email            string
	Link             string
	Extra            string
}

type CommentQuery struct {
	Id               string
	CreationTime     time.Time
	ModificationTime time.Time
	Content          string
	Uri              string
	Author           string
	Email            string
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

## Status

![](https://buildstats.info/github/chart/StardustDL/loment?branch=master)