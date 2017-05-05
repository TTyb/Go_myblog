{{define "headers"}}

<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <a type="button" class="navbar-toggle" data-toggle = "collapse"  data-target = "#target-menu">
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </a>
            <a class="navbar-brand brand" href="/">TTyb|个人网站</a>
        </div>

        <div style="font-family: serif;text-decoration: none;" class="collapse navbar-collapse navbar-responsive-collapse" id = "target-menu">
            <ul style="margin: 0 0 0;">
                <ul class="nav navbar-nav pull-right" style="margin: 0 0 0;">
                    <li>
                        <a data-toggle="dropdown" class="dropdown-toggle navbar-brand" href="javascript:void(0)">友情链接
                            <strong class="caret"></strong></a>
                        <ul class="dropdown-menu container-fluid" style="border-radius: 6px 6px 6px 6px;">
							<li><a href="http://www.cnblogs.com/TTyb">TTyb博客园</a></li>
							<li><a href="http://www.cnblogs.com/nima">一只尼玛博客园</a></li>
							<li><a href="https://github.com/TTyb">TTybGithub</a></li>
							<li><a href="http://www.cjhug.me/">搬砖的陈大师</a></li>
							<li><a href="http://www.cnblogs.com/starwater/">星水博客园</a></li>
						</ul>
                    </li>
                    {{if .IsLogin}}
                    <li><a class="navbar-brand" href="/login?exit=true">退出登录</a></li>
                    <li>
                        <a data-toggle="dropdown" class="dropdown-toggle navbar-brand" href="javascript:void(0)">管理<strong
                                class="caret"></strong></a>
                        <ul class="dropdown-menu container-fluid" style="border-radius: 6px 6px 6px 6px;">
                            <li><a href="/manage/topic">管理文章</a></li>
                            <li><a href="/manage/log">管理日志</a></li>
                        </ul>
                    </li>
                    {{else}}
                    <li><a class="navbar-brand" href="/login">管理员登录</a></li>
                    {{end}}
                </ul>
                <ul class="nav navbar-nav" style="margin: 0 0 0;">
                    <li {{if .IsHome}} class="active" {{end}}><a class="navbar-brand" href="/">首页</a></li>
                    <li {{if .IsTopic}} class="active" {{end}}><a class="navbar-brand" href="/topic">文章</a></li>
                    <li {{if .IsCategory}} class="active" {{end}}><a class="navbar-brand" href="/category">分类</a></li>
                    <li {{if .IsLog}} class="active" {{end}}><a class="navbar-brand" href="/log">日志</a></li>
                    <li {{if .IsLog}} class="active" {{end}}><a class="navbar-brand" href="/log">留言</a></li>
                    <li {{if .IsLog}} class="active" {{end}}><a class="navbar-brand" href="/log">联系</a></li>
                </ul>
            </ul>
        </div>
    </div>
</div>