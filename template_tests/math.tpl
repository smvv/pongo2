# Identity
{{number}}
{{var}}

# Arithmetic
{{var * 2}}
{{var / 2}}
{{(var * 2) / 2}}
{{var + 2}}
{{(var - 2) + 2}}
{{'33' + var}}

# Floating point issues
{{var * 25 * (1/100)}}

# Conditionals
{% if var > '4' %}yes{% else %}no{% endif %}
{% if var >= '4' %}yes{% else %}no{% endif %}
{% if var == '4' %}yes{% else %}no{% endif %}
{% if var != '4' %}yes{% else %}no{% endif %}
{% if var < '4' %}yes{% else %}no{% endif %}
{% if var <= '4' %}yes{% else %}no{% endif %}
