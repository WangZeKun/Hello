{{range $_,$a := .activities}}
<div class="zhenwen">
  <h2>{{$a.Name}}</h2>
  <xmp>{{$a.Introduction}} </xmp>
  <form action="/activities" method="post" target="_self">
    {{range $key, $val := $a.GetMessage}}
      {{$val}}:<input name="{{$val}}">
      {{end}} 
    <button class="button button-raised button-action button-pill" name="activity_id" value="{{.Id}}">点击报名</button>
  </form>
</div>
{{end}} <br>
<br>
<br>
<br>
<br>
<br>