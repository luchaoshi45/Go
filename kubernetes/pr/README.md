
# 如何提交PR到GitHub

### Part 1 - Local Repo Config 

1. 先 Fork 感兴趣项目，即 `kubeedge/examples`
2. Clone 到本地，`git clone https://github.com/luchaoshi45/examples.git`
3. 添加源项目 `kubeedge/examples` 作为 `upstream` 源，`git remote add upstream https://github.com/kubeedge/examples.git`
4. 禁止直接向 `upstream` 源 push，因为我们不是 kubeedge 的人，没有 push 的权限，要提交代码必须通过 Pull Request，`git remote set-url --push upstream no_push`
3. 创建并切换到本地的新分支 `sound-equipment-fault-detection`，`git checkout -b sound-equipment-fault-detection`
    + 本地分支 `master` 的作用是与远程 `upstream` 的最新代码同步
    + 本地分支 `sound-equipment-fault-detection` 则是我们修改代码的战场

### Part 2 - Fix Bug 

1. fix bug
2. 在当前 `sound-equipment-fault-detection` 分支提交本地修改，`git commit -m "XXXXXXX"`
3. Check `upstream` 源的最新状态
    + 在本地将 `upstream` 源的代码更新到最新，`git fetch upstream`
    + 将本地当前分支切换成 `master`，`git checkout master`
    + 将 `upstream/master` 的代码与本地当前默认分支，也就是本地 `master` 分支的代码融合：`git merge upstream/master`
4. 将本地分支 `sound-equipment-fault-detection` 上的修改融合到最新的 `master` 分支上
    + 将本地当前分支切换成 `sound-equipment-fault-detection`，`git checkout sound-equipment-fault-detection`
    + **将 `sound-equipment-fault-detection` 的代码置于已经更新到最新的 `master` 分支的最新代码之后**：`git rebase master`，如下图所示：
        <img src="https://www.atlassian.com/dam/jcr:5b153a22-38be-40d0-aec8-5f2fffc771e5/03.svg" width="600px">

5. 向 Github 上自己的 fork 项目 `kubeedge/examples` 的分支 `origin` 提交自己的修改，因为 Pull Request 是将两个 Github 上的 repo 比较，所以一定要将本地的修改先推送到自己的 fork repo 上，`git push origin sound-equipment-fault-detection:sound-equipment-fault-detection`，参看 Ref. 4
   


### Part 3 - Pull Request 

1. 提交 issue：描述发现的 Bug，这个可以作为对自己后面 Pull Request 的描述
2. 是在自己 Fork 的项目界面，即 kubeedge/examples 的 Pull requests 的 Tab 中点击 New pull request，后面会自动跳到 dmlc/gluon-cv 的界面，如下所示：

    <img src="https://raw.githubusercontent.com/YimianDai/images/master/Pull-PR-Empty.png" width="700px">

3. 解决 DCO（开发者证书）问题
<br> `git rebase HEAD~1 --signoff`
<br> `git push --force-with-lease origin sound-equipment-fault-detection`
## Reference

查看当前是哪个分支？在工作目录下, 
```shell
cat .git/HEAD
```

1. https://gist.github.com/YimianDai/7dcf6340fc435323a328634df0666f5e
2. 关于 `git rebase` 非常好的一篇文章，[Merging vs. Rebasing](https://www.atlassian.com/git/tutorials/merging-vs-rebasing)
3. [向 github 的开源项目提交 PR 的步骤](https://blog.csdn.net/u010857876/article/details/79035876)
4. [GIT: PUSHING TO A REMOTE BRANCH WITH A DIFFERENT NAME](https://penandpants.com/2013/02/07/git-pushing-to-a-remote-branch-with-a-different-name/)
5. `git push <REMOTENAME> <BRANCHNAME>`
