using Loment;
using Loment.Models;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using System;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Test.Base
{
    [TestClass]
    public class CommentsTest
    {
        LomentService Service { get; set; }

        [TestInitialize]
        public void Setup()
        {
            var client = Utils.CreateTestClient();

            Service = new LomentService(client);
        }

        [TestMethod]
        public async Task CreateGetAndDelete()
        {
            var comment = new Comment
            {
                Content = "abc"
            };
            var id = await Service.Create(comment);
            Assert.IsNotNull(id);

            var res = await Service.Get(id);

            Assert.AreEqual(comment.Content, res.Content);

            {
                var updated = await Service.Update(new Comment
                {
                    ID = id,
                    Content = "abcd"
                });
                Assert.IsTrue(updated);

                var res2 = await Service.Get(id);

                Assert.AreEqual("abcd", res2.Content);
            }

            var del = await Service.Delete(id);
            Assert.IsTrue(del);
        }

        [TestMethod]
        public async Task Query()
        {
            var items = await Service.Query(new CommentQuery());
            Assert.IsNotNull(items);
        }

        [TestCleanup]
        public void Clean()
        {
            Service.Client.Dispose();
        }
    }
}
