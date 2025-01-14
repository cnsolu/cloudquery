service          = "cloudflare"
output_directory = "."
add_generate     = true

resource "cloudflare" "" "waf" {
  path = "github.com/cloudflare/cloudflare-go/.WAFPackage"

  multiplex "CFZone" {
    path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ZoneMultiplex"
  }

  options {
    primary_keys = [
      "id"
    ]
  }

  deleteFilter "DeleteAccountZoneFilter" {
    path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.DeleteAccountZoneFilter"
  }

  userDefinedColumn "account_id" {
    description = "The Account ID of the resource."
    type        = "string"
    resolver "resolveCFAccount" {
      path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ResolveAccountId"
    }
  }

  column "id" {
    description = "The unique identifier of a WAF package."
  }

  column "name" {
    description = "The name of the WAF package."
  }

  column "description" {
    description = "A summary of the purpose/function of the WAF package."
  }

  column "zone_id" {
    description = "Zone identifier tag."
  }

  column "detection_mode" {
    description = "When a WAF package uses anomaly detection, each rule is given a score when triggered. If the total score of all triggered rules exceeds the sensitivity defined on the WAF package, the action defined on the package will be taken."
  }

  column "sensitivity" {
    description = "The sensitivity of the WAF package"
  }

  column "action_mode" {
    description = "The default action performed by the rules in the WAF package."
  }

  user_relation "cloudflare" "" "waf_rule_groups" {
    path = "github.com/cloudflare/cloudflare-go/.WAFGroup"

    userDefinedColumn "account_id" {
      description = "The Account ID of the resource."
      type        = "string"
      resolver "resolveCFAccount" {
        path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ResolveAccountId"
      }
    }

    userDefinedColumn "zone_id" {
      description = "Zone identifier tag."
      type        = "string"
      resolver "resolveCFZone" {
        path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ResolveZoneId"
      }
    }

    column "id" {
      description = "The unique identifier of the rule group."
    }

    column "name" {
      description = "The name of the rule group."
    }

    column "description" {
      description = "An informative summary of what the rule group does."
    }

    column "rules_count" {
      description = "The number of rules in the current rule group."
    }

    column "modified_rules_count" {
      description = "The number of rules within the group that have been modified from their default configuration."
    }

    column "package_id" {
      description = "The unique identifier of a WAF package."
    }

    column "mode" {
      description = "The state of the rules contained in the rule group. When on, the rules in the group are configurable/usable."
    }

    column "allowed_modes" {
      description = "The available states for the rule group."
    }

  }

  user_relation "cloudflare" "" "waf_rules" {
    path = "github.com/cloudflare/cloudflare-go/.WAFRule"

    userDefinedColumn "account_id" {
      description = "The Account ID of the resource."
      type        = "string"
      resolver "resolveCFAccount" {
        path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ResolveAccountId"
      }
    }

    userDefinedColumn "zone_id" {
      description = "Zone identifier tag."
      type        = "string"
      resolver "resolveCFZone" {
        path = "github.com/cloudquery/cloudquery/plugins/source/cloudflare/client.ResolveZoneId"
      }
    }

    column "id" {
      description = "The unique identifier of the WAF rule."
    }

    column "description" {
      description = "The public description of the WAF rule."
    }

    column "priority" {
      description = "The order in which the individual WAF rule is executed within its rule group."
    }

    column "group" {
      skip = true
    }

    userDefinedColumn "group" {
      description = "The rule group to which the current WAF rule belongs."
      type        = "JSON"
    }

    column "default_mode" {
      description = "The default action/mode of a rule."
    }

    column "mode" {
      description = "The action that the current WAF rule will perform when triggered. Applies to traditional (deny) WAF rules."
    }

    column "allowed_modes" {
      description = "The list of possible actions of the WAF rule when it is triggered."
    }
  }
}

resource "cloudflare" "" "waf_overrides" {
  path = "github.com/cloudflare/cloudflare-go/.WAFOverride"

  multiplex "CFZone" {
    path   = "github.com/cloudquery/cq-provider-cloudflare/client.ZoneMultiplex"
  }

  options {
    primary_keys = [
      "id"
    ]
  }

  deleteFilter "DeleteAccountZoneFilter" {
    path = "github.com/cloudquery/cq-provider-cloudflare/client.DeleteAccountZoneFilter"
  }

  userDefinedColumn "account_id" {
    description = "The Account ID of the resource."
    type        = "string"
    resolver "resolveCFAccount" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveAccountId"
    }
  }

  userDefinedColumn "zone_id" {
    description = "The Zone ID of the resource."
    type        = "string"
    resolver "resolveCFZone" {
      path = "github.com/cloudquery/cq-provider-cloudflare/client.ResolveZoneId"
    }
  }

  column "id" {
    description = "The unique identifier of the WAF override."
  }

  column "paused" {
    description = "When true, indicates that the WAF package is currently paused."
  }

  column "description" {
    description = "An informative summary of the current URI-based WAF override."
  }

  column "urls" {
    description = "The URLs to include in the current WAF override. You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns."
  }

  column "priority" {
    description = "The relative priority of the current URI-based WAF override when multiple overrides match a single URL. A lower number indicates higher priority. Higher priority overrides may overwrite values set by lower priority overrides."
  }

  column "groups" {
    description = "An object that allows you to enable or disable WAF rule groups for the current WAF override. Each key of this object must be the ID of a WAF rule group, and each value must be a valid WAF action (usually default or disable). When creating a new URI-based WAF override, you must provide a groups object or a rules object."
  }

  column "rules" {
    description = "The default action performed by the rules in the WAF package."
  }

  column "rewrite_action" {
    description = "Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object."
  }
}
