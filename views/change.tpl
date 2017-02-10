<br>
<br>
<br>
<br>
<form action="/change" method="POST" target="_self" style="text-align: center">
<div>原密码：&nbsp;&nbsp;&nbsp;&nbsp;
  <input name="password_former" type="password" class="ur" />
</div>
<div>新密码：&nbsp;&nbsp;&nbsp;&nbsp;
  <input name="password" type="password" class="ur" />
</div>
<div>重复密码：
  <input name="password_repeat" type="password" class="ur" />
</div>
<div style="color: red;font-size:12px"> {{.error}} </div>
<div>
  <input type="submit" value="&nbsp;&nbsp;修改 &nbsp;&nbsp;" class="button button-raised button-primary button-pill">
</div>
</form>
<br>
<br>
<br>
<br>