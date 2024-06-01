---
layout: base
title: index Page
eleventyImport:
  collections: ["post"]
---

{% block head %}

<link rel="stylesheet" href="/styles/index.css">
{% endblock %}

# Test

- Test1
- Test2

<ul>
{%- for post in collections.post | reverse -%}
  <li><a href="{{ post.url }}">{{ post.data.title  }} {{ page.date.toUTCString() }} </a></li>
{%- endfor -%}
</ul>
