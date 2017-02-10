<h2 style="padding: 20px; color:tomato">个人信息:</h2>
<nav style="display: flex">
<div class="text">
  <ul>
    <li>学号：{{.stu.Id}}</li>
    <li>性别：{{.stu.Gender}}</li>
    <li>年级：{{.stu.Grade}}</li>
  </ul>
</div>
<div class="text">
  <ul>
    <li>姓名：{{.stu.Name}}</li>
    <li>学段：{{.stu.Section}}</li>
    <li>班级：{{.stu.Class}}</li>
  </ul>
</div>
</nav>
<form action="/main" method="POST" target="_self">
  <nav style="display: flex">
<div class="text">
  <ul>
    <li>电话号码：
      <input name="telephone" value="{{.stu.Telephone}}" />
    </li>
    <li>微信号：
      <input name="weChat" value="{{.stu.WeChat}}" />
    </li>
  </ul>
</div>
<div class="text">
  <ul>
    <li> QQ：
      <input name="QQ" value="{{.stu.Qq}}" />
    </li>
  </ul>
</div>
</nav>
<div class="text" style="width: auto">
<ul>
  <li style="display: flex">个人简介：
<textarea name="jianjie" wrap="soft" cols="75" rows="5">{{.stu.Jianjie}}</textarea>
</li>
</ul>
</div>
<center>
  <div>
    <input type="submit" value="修改" class="button button-raised button-action button-pill xg">
  </div>
</center>
</form>
<h3 style="padding: 20px; color:tomato">活动综述：</h3>
<div style="align-items: center;display:inline-table;">
<table class="zs" cellspacing="10px">
  {{if not .Canjia}}
    <td>您还没有参加活动！！！</td>
  {{end}}
  {{range $_,$C := .Canjia}}
  <tr>
    <td>{{$C.GetTime}} </td>
    <td>{{$C.CheckActivity}}</td>
    <td>{{$C.Status}}</td>
  </tr>
  {{end}}
</table>
</div>
<br>
<br>
<br>