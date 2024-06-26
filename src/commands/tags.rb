# frozen_string_literal: true

require 'thor'
require 'dotenv/load'
require_relative '../config/github'
require_relative '../utils/log'
require_relative '../utils/output'
require_relative '../options/common'
require_relative '../utils/labels'
require_relative '../utils/tags'

class Tags < Thor
  desc 'list', 'List tags for the specified repo'
  option(*CommonOptions.repo)
  def list
    tags = GitHub.octokit.tags(options.repo)
    Utils::Tags.list(tags)
  end

  desc 'count', 'Count tags for the specified repo'
  option(*CommonOptions.repo)
  def count
    Utils::Tags.count(options.repo)
  end

  desc 'remove-stale', 'Remove old tags from the specified repo'
  option(*CommonOptions.repo)
  option(:limit, :type => :numeric, :default => 100, :desc => 'Quantity of tags to preserve')
  def remove_stale
    count
    if GitHub.octokit.tags(options.repo).size > options.limit
      Utils::Tags.list_after_limit(options.repo, options.limit)
      Utils::Tags.delete(Utils::Tags.after_limit(options.repo, options.limit))
    else
      puts "Tags count does not exceed specified limit (#{options.limit})."
    end
  end
end
